import { readFileSync, existsSync } from 'fs';
import { join } from 'path';
import { subscribe } from './subscribe';
import { AmqpService } from './common/amqp.service';
import { ElasticService } from './common/elastic.service';

const configPath = join(__dirname, 'config.json');
if (!existsSync(configPath)) {
  console.log('must set config.json');
  process.exit(1);
}
const config: any = JSON.parse(readFileSync(configPath, {
  encoding: 'utf8',
}));

AmqpService.registered(config.amqp).then(channel => {
  subscribe(
    channel,
    ElasticService.create(config.elastic).getClient(),
  );
}).catch(err => {
  if (err) {
    console.log(err);
    process.exit(1);
  }
});

