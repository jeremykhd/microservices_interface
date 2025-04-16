const amqp = require('amqplib');

let channel = null;

const connectToRabbitMQ = async () => {
  try {
    const connection = await amqp.connect(process.env.RABBITMQ_URI || 'amqp://localhost');
    channel = await connection.createChannel();
    
    // Declare exchanges
    await channel.assertExchange('workout-events', 'topic', { durable: true });
    
    console.log('Connected to RabbitMQ');
  } catch (error) {
    console.error('Error connecting to RabbitMQ:', error);
    throw error;
  }
};

const publishMessage = async (routingKey, message) => {
  if (!channel) {
    throw new Error('RabbitMQ channel not initialized');
  }

  try {
    await channel.publish(
      'workout-events',
      routingKey,
      Buffer.from(JSON.stringify(message)),
      { persistent: true }
    );
    console.log(`Message published with routing key: ${routingKey}`);
  } catch (error) {
    console.error('Error publishing message:', error);
    throw error;
  }
};

module.exports = {
  connectToRabbitMQ,
  publishMessage
}; 