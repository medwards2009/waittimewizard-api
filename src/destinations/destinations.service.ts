import { Injectable } from "@nestjs/common";

import { ApiClientService } from "../common/services/apiclient.service";
import { DestinationDto } from "./dto/destination.dto";

@Injectable()
export class DestinationsService {
    constructor(private readonly apiClientService: ApiClientService) {}
    
    async getDestinations(): Promise<DestinationDto[]> {
        const response = await this.apiClientService.get('/destinations');
        return response.data as DestinationDto[];
    }
}