import { Module } from '@nestjs/common';
import { LiveService } from './live.service';
import { LiveController } from './live.controller';
import { ApiClientService } from 'src/common/services/apiclient.service';

@Module({
  controllers: [LiveController],
  providers: [LiveService],
  imports: [ApiClientService]
})
export class LiveModule {}