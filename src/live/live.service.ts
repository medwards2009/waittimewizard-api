import { Injectable } from "@nestjs/common";

import { ApiClientService } from "../common/services/apiclient.service";
import { LiveDto } from "./dto/live.dto";

@Injectable()
export class LiveService {
    constructor(private readonly apiClientService: ApiClientService) {}
    
    async getLive(id: string): Promise<LiveDto[]> {
        const response = await this.apiClientService.get(`/entity/${id}/live`);
        return response.data as LiveDto[];
    }
}