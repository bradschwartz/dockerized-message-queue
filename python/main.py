#!/usr/bin/env python
import sys, time
import logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)
import pika


def get_conn(server=None):
  if server is None:
    server = 'localhost'

  max_tries = 5
  while True and max_tries != 0:
    try:
      connection = pika.BlockingConnection(pika.ConnectionParameters(server))
      logger.info(f"Connected to {server}")
      return connection
    except Exception as err:
      logger.error(f"Failed to connect to {server}")
      logger.error(err)
      time.sleep(30)
      max_tries -= 1
  else:
    logger.critical("Exceeded max tries alloted. Try again later")
    sys.exit(1)

def callback(ch, method, properties, body):
  logger.info(f" [x] Received {body}")


def main():
  server = 'mqserver'
  queue = 'worker'
  connection = get_conn(server)
  logger.info(connection)
  channel = connection.channel()
  logger.info(channel)

  channel.queue_declare(queue)
  channel.basic_consume(queue=queue,
                        auto_ack=True,
                        on_message_callback=callback)
  channel.start_consuming()

  connection.close()

if __name__ == '__main__':
  main()