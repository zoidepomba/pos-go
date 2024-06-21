import boto3

# Nomes dos serviços
services = [
    "prd-avi-chatbot-sales-service"
]

# Nome do cluster ECS
cluster_name = "prd-avi"

# Cria um cliente ECS
ecs_client = boto3.client('ecs')

# Função para escalar um serviço para 2 instâncias
def scale_service_to_2_instances(service_name):
    try:
        response = ecs_client.update_service(
            cluster=cluster_name,
            service=service_name,
            desiredCount=20
        )
        print(f"Scaled up {service_name} to 2 instances")
        print(response)
    except Exception as e:
        print(f"Failed to scale up {service_name}: {str(e)}")

# Loop para percorrer a lista de serviços e escalar cada um
for service in services:
    scale_service_to_2_instances(service)