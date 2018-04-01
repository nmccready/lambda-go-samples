package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const msgTemplate = `{ "text": "Instances",
  "attachments": [
  {{ range . }}
  {{ range .Instances }}
  { "title" : "{{- range .Tags }}
      {{- if eq ( Deref .Key ) "Name" }}
        {{- .Value}}
      {{- end}}
    {{- end}}",
    
    
    "footer": "{{ .LaunchTime }}",
    {{ if eq "running" (Deref .State.Name) }} "color":"#36a64f" ,{{end}}
    "fields": [
       {
          "title": "PublicIP",
          "value": "{{ .PublicIpAddress}}",
          "short": true
        },
       {
          "title": "Zone",
          "value": "{{ .Placement.AvailabilityZone}}",
          "short": true
        },
       {
          "title": "State",
          "value": "{{ .State.Name }}",
          "short": true
        },
       {
          "title": "InstanceId",
          "value": "{{ .InstanceId}}",
          "short": true
        }
    ]
    },
    {{end}}
    {{end}}
  ]
}`
const asciiTemplate = `{ "text": " instances: ` + "```" + `
-------------------------------------------------------------------------------------------------------------+
{{ printf "%-20s" "instanceId" }} | 
{{- printf " %-32s" "name" }} | 
{{- printf " %-16s" "publicIp" }} | 
{{- printf " %-12s" "region" }} | 
{{- printf " %-16s" "started" }} | 
-------------------------------------------------------------------------------------------------------------+
{{ range . }}
  {{- range .Instances }}
    {{- .InstanceId | Deref | printf "%-20s" }} |
    {{- range .Tags }}
      {{- if eq ( Deref .Key ) "Name" }}
        {{- .Value | Deref | printf " %-32s"}}
      {{- end}}
    {{- end}} |
   {{- .PublicIpAddress | Deref | printf " %-16s" }} |
   {{- .Placement.AvailabilityZone | Deref | printf " %-12s" }} | 
   {{- .LaunchTime | Hours  | printf "%-16s"}} | 
{{ end}}
{{- end}}-------------------------------------------------------------------------------------------------------------+` +
	"```" + `"}`

const asciiTemplateHours = `{ "text": " instances: ` + "```" + `
----------------------------------------------------------------------------------------------------------------+
{{ printf "%-20s" "instanceId" }} | 
{{- printf " %-30s" "name" }} | 
{{- printf " %-14s" "publicIp" }} | 
{{- printf " %-12s" "region" }} | 
{{- printf " %-5s" "hours" }} | 
{{- printf " %-9s" "insType" }} | 
{{- printf " %-3s" "spt" }} | 
----------------------------------------------------------------------------------------------------------------+
{{ range . }}
  {{- range .Instances }}
    {{- .InstanceId | Deref | printf "%-20s" }} |
    {{- range .Tags }}
      {{- if eq ( Deref .Key ) "Name" }}
        {{- .Value | Deref | printf " %-30s"}}
      {{- end}}
    {{- end}} |
   {{- .PublicIpAddress | Deref | printf " %-14s" }} |
   {{- .Placement.AvailabilityZone | Deref | printf " %-12s" }} |
   {{- .LaunchTime | Hours | printf "%5sh" }} | 
   {{- .InstanceType | Deref| printf " %-9s"}} |
   {{- if .InstanceLifecycle }}  X {{else}}    {{end}} |
{{ end}}
{{- end}}----------------------------------------------------------------------------------------------------------------+` +
	"```" + `"}`

/*
 */

func formatInstances(reservations []*ec2.Reservation, toAscii bool) string {
	actualTempl := msgTemplate
	if toAscii {
		actualTempl = asciiTemplateHours
	}
	tmpl, err := template.New("test").Funcs(template.FuncMap{
		"Deref": func(i *string) string {
			if i != nil {
				return *i
			} else {
				return "nil"
			}
		},
		"Hours": func(t *time.Time) string {
			hour := strings.Split(fmt.Sprint(time.Since(*t).Truncate(time.Hour)), "h")[0]
			return hour
		},
	}).Parse(actualTempl)

	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, reservations)
	if err != nil {
		panic(err)
	}
	return b.String()
}

func awsInstancesInRegion(reg string) []*ec2.Reservation {

	fmt.Println("query reg:", reg, "started ...")
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(reg),
	}))
	svc := ec2.New(sess)
	din, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name:   aws.String("instance-state-name"),
				Values: aws.StringSlice([]string{"running"}),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return din.Reservations
}

func awsInsatncesMsg(respUrl string, toAscii bool) {

	_, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic(err)
	}

	if os.Getenv("regions") == "" {
		os.Setenv("regions", "eu-west-1,eu-central-1,us-east-1")
	}
	regions := strings.Split(os.Getenv("regions"), ",")
	chIns := make(chan []*ec2.Reservation, len(regions))

	for _, r := range regions {
		go func(reg string) {
			chIns <- awsInstancesInRegion(reg)
		}(r)
	}

	allInstances := make([]*ec2.Reservation, 0)
	for range regions {
		allInstances = append(allInstances, <-chIns...)
	}

	msg := formatInstances(allInstances, toAscii)

	resp, err := http.Post(respUrl, "application/json", bytes.NewBufferString(msg))
	if err != nil {
		panic(err)
	}
	fmt.Println("response:", resp.Status)
}

/*
func main() {
	awsInsatncesMsg(os.Getenv("respUrl"), true)
}
*/
