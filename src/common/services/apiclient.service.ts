import { Injectable } from '@nestjs/common';
import axios, { AxiosInstance } from 'axios';

@Injectable()
export class ApiClientService {
  private client: AxiosInstance;
  
  constructor() {
    this.client = axios.create({
      baseURL: 'https://api.themeparks.wiki/v1'
    });
  }

  async get(url: string) {
    return this.client.get(url);
  }
}