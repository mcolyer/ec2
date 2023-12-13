package main

import (
	"fmt"
	"os"

	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: <program> <region> <start|stop> <instance-id>")
		os.Exit(1)
	}

	region := os.Args[1]
	action := os.Args[2]
	instanceId := os.Args[3]

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	svc := ec2.New(sess)

	switch action {
	case "start":
		startInstance(svc, instanceId)
	case "stop":
		stopInstance(svc, instanceId)
	default:
		fmt.Println("Invalid action. Use 'start' or 'stop'.")
	}
}
func startInstance(svc ec2iface.EC2API, instanceId string) {
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceId),
		},
	}

	_, err := svc.StartInstances(input)
	if err != nil {
		fmt.Println("Error starting instance:", err)
		return
	}

	fmt.Println("Instance is starting, waiting for IP address...")
	ip, err := waitForIp(svc, instanceId)
	if err != nil {
		fmt.Println("Error getting IP address:", err)
		return
	}

	fmt.Printf("Instance started. IP Address: %s\n", ip)
}

func waitForIp(svc ec2iface.EC2API, instanceId string) (string, error) {
	for {
		descInput := &ec2.DescribeInstancesInput{
			InstanceIds: []*string{aws.String(instanceId)},
		}
		result, err := svc.DescribeInstances(descInput)
		if err != nil {
			return "", err
		}

		for _, r := range result.Reservations {
			for _, i := range r.Instances {
				if *i.InstanceId == instanceId && *i.State.Name == ec2.InstanceStateNameRunning {
					return *i.PublicIpAddress, nil
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func stopInstance(svc *ec2.EC2, instanceId string) {
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceId),
		},
	}

	result, err := svc.StopInstances(input)
	if err != nil {
		fmt.Println("Error stopping instance:", err)
		return
	}

	fmt.Println("Instance stopped:", result.StoppingInstances)
}
