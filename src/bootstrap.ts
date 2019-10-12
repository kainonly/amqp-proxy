import { Client } from '@elastic/elasticsearch';

const client = new Client({ node: 'http://imac:9200' });
console.log(client);
