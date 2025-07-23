import { AxiosError, type AxiosInstance, type AxiosResponse } from 'axios';

export type ApiRequestMethod = 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE';

export interface ApiRequestHeaders {
  [x: string]: string | number | boolean;
}

export interface ApiRequestProps<S, F, B, P> {
  /* request to */
  instance: AxiosInstance;
  method?: ApiRequestMethod;
  path: string;
  /* common headers */
  token?: string;
  lang?: string;
  /* transport data */
  headers?: ApiRequestHeaders;
  params?: P;
  body?: B;
  /* response */
  success: (response: AxiosResponse<unknown, unknown>) => S;
  failure: (error: AxiosError) => F;
  /* configuration */
  abort?: AbortController;
  timeout?: number;
}

export async function apiRequest<S, F = undefined, B = undefined, P = undefined>({
  /* request to */
  instance,
  method = 'GET',
  path,
  /* common headers */
  token,
  lang,
  /* transport data */
  headers,
  params,
  body,
  /* response */
  success,
  failure,
  /* configuration */
  abort,
  timeout = 5 * 1000,
}: ApiRequestProps<S, F, B, P>): Promise<S | F> {
  const requestHeaders = {
    /* base headers */
    'Content-Type': 'application/json',
    ...(lang && { 'Accept-Language': lang }),
    /* authorization */
    ...(token && { Authorization: `Bearer ${token}` }),
    /* others */
    ...headers,
  };

  try {
    const response = await instance.request({
      headers: requestHeaders,
      method,
      url: path,
      params,
      data: body,
      signal: abort?.signal,
      timeout,
    });

    return success(response);
  } catch (error) {
    return failure(error as AxiosError);
  }
}
