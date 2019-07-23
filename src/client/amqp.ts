import { connect } from 'amqplib';

const AmqpChannel = async (uri: string) => {
  const conn = await connect(uri);
  return await conn.createChannel();
};

export { AmqpChannel };
