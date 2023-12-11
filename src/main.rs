use clap::{App, Arg, SubCommand};
use rusoto_core::{Region, HttpClient};
use rusoto_ec2::{Ec2, Ec2Client, StartInstancesRequest, StopInstancesRequest};
use rusoto_credential::EnvironmentProvider;

#[tokio::main]
async fn main() {
    let matches = App::new("EC2 Instance Manager")
        .version("1.0")
        .author("Your Name")
        .about("Manages EC2 instances")
        .subcommand(SubCommand::with_name("start")
            .about("Starts an EC2 instance")
            .arg(Arg::with_name("INSTANCE_ID")
                .help("The ID of the instance to start")
                .required(true)
                .index(1)))
        .subcommand(SubCommand::with_name("stop")
            .about("Stops an EC2 instance")
            .arg(Arg::with_name("INSTANCE_ID")
                .help("The ID of the instance to stop")
                .required(true)
                .index(1)))
        .get_matches();

    // Create EC2 client
    let client = Ec2Client::new_with(
        HttpClient::new().expect("failed to create request dispatcher"),
        EnvironmentProvider::default(),
        Region::default(),
    );

	match matches.subcommand() {
		Some(("start", sub_m)) => {
			let instance_id = sub_m.value_of("INSTANCE_ID").unwrap();
			start_instance(&client, instance_id).await;
		},
		Some(("stop", sub_m)) => {
			let instance_id = sub_m.value_of("INSTANCE_ID").unwrap();
			stop_instance(&client, instance_id).await;
		},
		_ => eprintln!("Invalid command or no subcommand was provided"),
	}
}

async fn start_instance(client: &Ec2Client, instance_id: &str) {
    let start_request = StartInstancesRequest {
        instance_ids: vec![instance_id.to_string()],
        ..Default::default()
    };

    match client.start_instances(start_request).await {
        Ok(output) => println!("Instance started: {:?}", output),
        Err(error) => eprintln!("Error starting instance: {}", error),
    }
}

async fn stop_instance(client: &Ec2Client, instance_id: &str) {
    let stop_request = StopInstancesRequest {
        instance_ids: vec![instance_id.to_string()],
        ..Default::default()
    };

    match client.stop_instances(stop_request).await {
        Ok(output) => println!("Instance stopped: {:?}", output),
        Err(error) => eprintln!("Error stopping instance: {}", error),
    }
}
