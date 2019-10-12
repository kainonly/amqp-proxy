import { Channel } from 'amqplib/callback_api';
import { Client } from '@elastic/elasticsearch';
import * as Ajv from 'ajv';

export function subscribe(channel: Channel, elastic: Client, rule: any) {
  channel.assertExchange('elastic.stash.exchange', 'direct', {
    durable: true,
  });
  channel.assertQueue('elastic.stash.basic', {
    durable: true,
  });
  channel.bindQueue('elastic.stash.basic', 'elastic.stash.exchange', '/');
  channel.consume('elastic.stash.basic', async (msg) => {
    try {
      const ajv = new Ajv();
      const isJson = ajv.validate({
        type: 'string',
        contentMediaType: 'application/json',
      }, msg.content.toString());
      if (!isJson) {
        console.log(ajv.errors);
        channel.ack(msg);
        return;
      }
      const message = JSON.parse(msg.content.toString());
      const basic = ajv.validate({
        type: 'object',
        required: ['index', 'body'],
        properties: {
          index: {
            type: 'string',
          },
          body: {
            type: 'object',
          },
        },
      }, message);
      if (!basic) {
        console.log(ajv.errors);
        channel.ack(msg);
        return;
      }
      if (!rule.hasOwnProperty(message.index)) {
        console.log('The index json schema does not exist!');
        channel.ack(msg);
        return;
      }
      const special = ajv.validate(rule[message.index], message.body);
      if (!special) {
        console.log(ajv.errors);
        channel.ack(msg);
        return;
      }
      const response = await elastic.index(message);
      if (response.statusCode === 201) {
        console.log('stash success');
        channel.ack(msg);
      }
    } catch (e) {
      console.log(e.message);
      channel.ack(msg);
    }
  }, {
    noAck: false,
  }, err => {
    console.log(err);
  });
}
