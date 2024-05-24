import json
import random
import time
from faker import Faker
from datetime import datetime, timezone
import pika

fake = Faker()

def generate_user():
    genre = random.choice(['M', 'F'])
    first_name = fake.first_name_male() if genre == "M" else fake.first_name_female()
    last_name = fake.last_name()
    
    return {
        "id": fake.uuid4(),
        "first_name": first_name,
        "last_name": last_name,
        "email": f"{first_name.lower()}.{last_name.lower()}@{fake.domain_name()}",
        # "created_at": datetime.now(timezone.utc).strftime('%Y-%m-%d')
        # "created_at": datetime.now().strftime('%Y-%m-%d')
    }

def main():
    connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
    channel = connection.channel()

    # Declare the DLX exchange
    channel.exchange_declare(exchange='dlx_exchange', exchange_type='direct', durable=True)

    # Declare the DLX queue
    channel.queue_declare(queue='dlx_queue', durable=True)

    # Bind DLX queue to DLX exchange
    channel.queue_bind(exchange='dlx_exchange', queue='dlx_queue', routing_key='dlx_routing_key')

    # Declare the main queue with DLX parameters
    args = {
        'x-dead-letter-exchange': 'dlx_exchange',
        'x-dead-letter-routing-key': 'dlx_routing_key'
    }

    channel.queue_declare(queue='user', durable=True, arguments=args)

    curr_time = datetime.now()

    # while (datetime.now() - curr_time).seconds < 120:
    for _ in range(10000):
        try:
            user = generate_user()

            channel.basic_publish(
                exchange='',
                routing_key='user',
                body=json.dumps(user)
            )

            print(f" [x] Sent {user}")

            # time.sleep(2)
        except Exception as e:
            print(e)
            time.sleep(1)

    connection.close()

if __name__ == "__main__":
    main()
