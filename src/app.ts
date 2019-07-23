import { Mongo } from './client/mongo';
import { AmqpChannel } from './client/amqp';

const Ajv = require('ajv');

const App = async (env: any) => {
  try {
    const db = (await Mongo(env.mongo_uri)).db(env.database);
    const channel = await AmqpChannel(env.amqp_uri);
    await channel.assertExchange(env.exchange, 'direct', { durable: true });
    await channel.assertQueue(env.queue, { durable: true });
    await channel.bindQueue(env.queue, env.exchange, '');
    return await channel.consume(env.queue, async (msg) => {
      const param = JSON.parse(msg.content.toString());
      const ajv = new Ajv();
      const validate = ajv.validate({
        required: ['appid', 'namespace', 'raws'],
        properties: {
          appid: {
            type: 'string',
          },
          namespace: {
            type: 'string',
          },
          raws: {
            type: 'object',
          },
        },
      }, param);

      if (!validate) {
        channel.ack(msg);
        console.error(ajv.errors);
        return;
      }

      const whitelist = await db.collection('whitelist').findOne({
        appid: param.appid,
      });

      if (whitelist.namespace.indexOf(param.namespace) === -1) {
        channel.ack(msg);
      } else {
        const collection = db.collection(param.namespace);
        collection.insertOne({
          appid: param.appid,
          raws: param.raws,
        }, (err) => {
          if (!err) channel.ack(msg);
        });
      }
    });
  } catch (e) {
    return e;
  }
};

export { App };
