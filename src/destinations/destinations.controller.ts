import { Controller, Get } from '@nestjs/common';
import { ApiOperation, ApiResponse, ApiTags } from '@nestjs/swagger';
import { DestinationsService } from './destinations.service';
import { DestinationDto } from './dto/destination.dto';

@ApiTags('destinations')
@Controller('destinations')
export class DestinationsController {
    constructor(private readonly destinationsService: DestinationsService) {}
    
    @Get()
    @ApiOperation({ summary: 'Get all theme park destinations' })
    @ApiResponse({ 
        status: 200, 
        description: 'List of all available theme park destinations',
        type: [DestinationDto]
    })
    async getAllDestinations(): Promise<DestinationDto[]> {
        return this.destinationsService.getDestinations();
    }
}
