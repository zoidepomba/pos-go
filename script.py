import boto3
import time
from datetime import datetime, timedelta

def lambda_handler(event, context):
    logs_client = boto3.client('logs')
    sns_client = boto3.client('sns')

    # Define the SNS topic ARN and the threshold for sending a notification
    sns_topic_arn = 'arn:aws:sns:us-east-1:571186639223:prd-avi-timeout-topic'
    notification_threshold = 3000 # Threshold defined when an identity has more than X occurrences within a 5-minute duration.

    # Define the query with a limit to get the top 10 results
    query = """fields @timestamp, @message, @logStream, @log
                | filter @logStream like "prd-avi-chatbot-corporate" 
                | filter @message like 'ecommresbff/v1/mind/addresses/viabilities'
                | sort @timestamp desc
                | limit 10000"""
    
    query2 = """fields @timestamp, @message, @logStream, @log
                | filter @logStream like "prd-avi-chatbot-corporate" 
                | filter @message like 'ecommresbff/v1/mind/addresses/viabilities'
                | sort @timestamp desc
                | limit 10000"""

    # Calculate the time window of the last 5 minutes
    end_time = int(datetime.now().timestamp())
    start_time = int((datetime.now() - timedelta(minutes=5)).timestamp())

    # Start the query
    start_query_response = logs_client.start_query(
        logGroupName='ecs-prd-avi',
        startTime=start_time,
        endTime=end_time,
        queryString=query,
    )
    query_id = start_query_response['queryId']
    
    #start the query 2
    start_query_response2 = logs_client.start_query(
        logGroupName='ecs-prd-avi',
        startTime=start_time,
        endTime=end_time,
        queryString=query2,
    )


    # Wait for the query to start running or complete
    time.sleep(10)

    response = logs_client.get_query_results(queryId=query_id)

    # Wait until the query status is 'Complete' or until a timeout is reached
    timeout = time.time() + 30
    while response['status'] in ['Scheduled', 'Running'] and time.time() < timeout:
        time.sleep(5)
        response = logs_client.get_query_results(queryId=query_id)


    if response['status'] == 'Complete':
        result_count = len(response['results'])

        if result_count > 1:
                    message = f"Alerta de Falhas de Autenticação:\n\nMais de um log encontrado com falha de autenticação dentro dos últimos 5 minutos.\n\nTotal de resultados: {result_count}\n"
                    message += "\nFavor verificar, pois pode estar relacionado a falhas de autenticação.\n"
                    #message += "\nPara mais detalhes: https://clarodigital.atlassian.net/wiki/spaces/EC/pages/386608922792/Alerta+de+Loop+-+Investiga+o+e+a+o"

                    sns_client.publish(
                        TopicArn=sns_topic_arn,
                        Message=message,
                        Subject='Alerta de Falhas de Autenticação'
                    )
    else:
        print(f"Query did not complete within the expected time. Status: {response['status']}")

    return {
        'statusCode': 200,
        'body': 'Query processing complete'
    }
