import { Client } from '@elastic/elasticsearch';

export class ElasticService {
  private client: Client;

  static create(options: any) {
    return new ElasticService(options);
  }

  constructor(options: any) {
    this.client = new Client(options);
  }

  getClient() {
    return this.client;
  }
}
