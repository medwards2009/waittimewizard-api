import { ApiClientService } from "./services/apiclient.service";

import { Module } from '@nestjs/common';

@Module({
  providers: [ApiClientService],  // Add this line to register the service
  exports: [ApiClientService]
})
export class CommonModule {}