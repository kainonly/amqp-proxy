import { App } from './app';
import { config } from 'dotenv';

const env = config().parsed;

App(env).then(result => {
  console.log(result);
});
