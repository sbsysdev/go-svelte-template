import axios from 'axios';

export const apiV1 = axios.create({
  baseURL: 'http://localhost:8080/api/v1',
  timeout: 5000, // 5 seconds
});
