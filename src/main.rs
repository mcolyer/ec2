use rusoto_core::{Region, HttpClient};
use rusoto_ec2::{Ec2, Ec2Client, StartInstancesRequest};
use rusoto_credential::EnvironmentProvider;

#[tokio::main]
async fn main() {
    // Create a client with credentials from environment variables
    let client = Ec2Client::new_with(
        HttpClient::new().expect("failed to create request dispatcher"),
        EnvironmentProvider::default(),
        Region::default(),
    );

    // Define your instance reservation ID here
    let reservation_id = "your-reservation-id";

    // Create the StartInstancesRequest
    let start_request = StartInstancesRequest {
        instance_ids: vec![reservation_id.to_string()],
        ..Default::default()
    };

    // Attempt to start the instance
    match client.start_instances(start_request).await {
        Ok(output) => {
            println!("Instance started: {:?}", output);
        }
        Err(error) => {
            eprintln!("Error starting instance: {}", error);
        }
    }
}
