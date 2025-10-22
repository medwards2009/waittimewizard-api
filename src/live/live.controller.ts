import { Controller, Get, Param } from '@nestjs/common';
import { ApiOperation, ApiResponse } from '@nestjs/swagger';

import { LiveService } from './live.service';
import { LiveDto } from './dto/live.dto';

@Controller('live')
export class LiveController {
    constructor(private readonly liveService: LiveService) {}

    @Get(':id')
    @ApiOperation({ summary: 'Get live theme park data' })
    @ApiResponse({ 
        status: 200, 
        description: 'List of all available live data for a specific destination',
        type: [LiveDto]
    })
    async getLiveById(@Param('id') id: string): Promise<LiveDto[]> {
        return this.liveService.getLive(id);
    }
}
