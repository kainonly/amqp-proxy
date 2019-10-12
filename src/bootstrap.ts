import { env } from 'process';
import { subscribe } from './subscribe';
import { AmqpService } from './common/amqp.service';
import { ElasticService } from './common/elastic.service';

AmqpService.registered(env).ready().then(channel => {
  const elastic = new ElasticService();

}).catch(err => {
  if (err) {
    console.log(err);
    process.exit(1);
  }
});

