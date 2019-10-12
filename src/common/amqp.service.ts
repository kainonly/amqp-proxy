import { Channel, connect } from 'amqplib/callback_api';

export class AmqpService {
  private options = {};

  static registered(options: any) {
    return new AmqpService(options).ready();
  }

  constructor(options: any) {
    this.options = options;
  }

  ready(): Promise<Channel> {
    return new Promise((resolve, reject) => {
      connect(this.options, (error, conn) => {
        if (error) {
          reject(error);
        }
        conn.createChannel((err, channel) => {
          if (err) {
            reject(err);
          } else {
            console.debug('SERVICE::SUBSCRIBE_START');
            resolve(channel);
          }
        });
      });
    });
  }
}
