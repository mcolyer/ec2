package main

import (
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
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

func startInstance(svc *ec2.EC2, instanceId string) {
    input := &ec2.StartInstancesInput{
        InstanceIds: []*string{
            aws.String(instanceId),
        },
    }

    result, err := svc.StartInstances(input)
    if err != nil {
        fmt.Println("Error starting instance:", err)
        return
    }

    fmt.Println("Instance started:", result.StartingInstances)
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
