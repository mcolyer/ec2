const AWS = require('aws-sdk');

const ec2 = new AWS.EC2({ region: 'us-west-1' });

const startInstance = async (instanceId) => {
    try {
        await ec2.startInstances({ InstanceIds: [instanceId] }).promise();
        console.log("Instance is starting...");

        // Poll for the instance state
        let describeParams = { InstanceIds: [instanceId] };
        let running = false;
        while (!running) {
            let data = await ec2.describeInstances(describeParams).promise();
            let state = data.Reservations[0].Instances[0].State.Name;
            if (state === 'running') {
                running = true;
                console.log("Instance started. IP Address:", data.Reservations[0].Instances[0].PublicIpAddress);
            } else {
                await new Promise(resolve => setTimeout(resolve, 10000)); // wait for 10 seconds
            }
        }
    } catch (error) {
        console.error("Error starting instance:", error);
    }
};

const stopInstance = async (instanceId) => {
    try {
        await ec2.stopInstances({ InstanceIds: [instanceId] }).promise();
        console.log("Instance is stopping...");
    } catch (error) {
        console.error("Error stopping instance:", error);
    }
};

const main = async () => {
    if (process.argv.length !== 4) {
        console.log("Usage: node main.js <start|stop> <instance-id>");
        return;
    }

    const command = process.argv[2];
    const instanceId = process.argv[3];

    if (command === "start") {
        await startInstance(instanceId);
    } else if (command === "stop") {
        await stopInstance(instanceId);
    } else {
        console.log("Invalid command. Use 'start' or 'stop'.");
    }
};

main();
