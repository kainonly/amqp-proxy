import { connect } from 'amqplib/callback_api';

export class AmqpService {
  private options = {};

  static registered(options: any) {
    return new AmqpService(options);
  }

  constructor(options: any) {
    this.setOptions(options);
  }

  private setOptions(options: any) {
    this.options = {
      hostname: options.amqp_hostname,
      username: options.amqp_username,
      password: options.amqp_password,
    };
    if (options.hasOwnProperty('protocol')) {
      Reflect.set(this.options, 'protocol', options.amqp_protocol);
    }
    if (options.hasOwnProperty('port')) {
      Reflect.set(this.options, 'port', parseInt(options.amqp_port));
    }
    if (options.hasOwnProperty('heartbeat')) {
      Reflect.set(this.options, 'heartbeat', parseInt(options.amqp_heartbeat));
    }
    if (options.hasOwnProperty('vhost')) {
      Reflect.set(this.options, 'vhost', options.amqp_vhost);
    }
  }

  ready() {
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
