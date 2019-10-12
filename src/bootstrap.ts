import { readFileSync, existsSync } from 'fs';
import { join } from 'path';
import * as Ajv from 'ajv';
import { subscribe } from './subscribe';
import { AmqpService } from './common/amqp.service';
import { ElasticService } from './common/elastic.service';

const configPath = join(__dirname, 'config.json');
if (!existsSync(configPath)) {
  console.log('Please set the config.json project configuration!');
  process.exit(1);
}
const config: any = JSON.parse(readFileSync(configPath, {
  encoding: 'utf8',
}));
const ajv = new Ajv();
const valid = ajv.validate({
  type: 'object',
  required: ['amqp', 'elastic', 'rule'],
  properties: {
    amqp: {
      type: 'object',
      required: ['hostname'],
      properties: {
        protocol: {
          enum: ['amqp', 'amqps'],
        },
        hostname: {
          type: 'string',
        },
        port: {
          type: 'number',
        },
        username: {
          type: 'string',
        },
        password: {
          type: 'string',
        },
        locale: {
          type: 'string',
        },
        frameMax: {
          type: 'number',
        },
        heartbeat: {
          type: 'number',
        },
        vhost: {
          type: 'string',
        },
      },
    },
    elastic: {
      type: 'object',
      properties: {
        node: {
          type: ['string', 'array', 'object'],
        },
        nodes: {
          type: ['string', 'array'],
        },
      },
    },
    rule: {
      type: 'object',
    },
  },
}, config);
if (!valid) {
  console.log(ajv.errors);
  process.exit(1);
}

AmqpService.registered(config.amqp).then(channel => {
  subscribe(
    channel,
    ElasticService.create(config.elastic).getClient(),
    config.rule,
  );
}).catch(err => {
  if (err) {
    console.log(err);
    process.exit(1);
  }
});

