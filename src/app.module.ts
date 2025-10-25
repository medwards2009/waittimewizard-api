import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { DestinationsController } from './destinations/destinations.controller';
import { LiveController } from './live/live.controller';
import { CommonModule } from './common/common.module';
import { DestinationsService } from './destinations/destinations.service';
import { LiveService } from './live/live.service';
import { RootController } from './root/root.controller';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true, // Make configuration available throughout the app
      envFilePath: '.env',
    }),
    CommonModule,
  ],
  controllers: [DestinationsController, LiveController, RootController],
  providers: [DestinationsService, LiveService],
})
export class AppModule {}
