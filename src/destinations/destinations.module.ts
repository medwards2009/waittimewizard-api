import { Module } from '@nestjs/common';
import { DestinationsService } from './destinations.service';
import { DestinationsController } from './destinations.controller';
import { ApiClientService } from 'src/common/services/apiclient.service';

@Module({
  controllers: [DestinationsController],
  providers: [DestinationsService],
  imports: [ApiClientService]
})
export class DestinationsModule {}