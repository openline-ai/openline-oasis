import {GraphQLClient} from "graphql-request";

var client: GraphQLClient;

export function setClient(cl: GraphQLClient): void {
  client = cl;
}

export function useGraphQLClient(): GraphQLClient {
  return client;
}