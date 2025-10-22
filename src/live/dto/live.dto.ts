import { ApiProperty } from '@nestjs/swagger';
import { Type } from 'class-transformer';
import { IsString, IsArray, ValidateNested, IsEnum, IsOptional, IsNumber, IsDateString } from 'class-validator';

class PriceDto {
    @ApiProperty({ description: 'Price amount' })
    @IsNumber()
    amount: number;

    @ApiProperty({ description: 'Currency code' })
    @IsString()
    currency: string;

    @ApiProperty({ description: 'Formatted price string' })
    @IsString()
    formatted: string;
}

class StandardQueueDto {
    @ApiProperty({ description: 'Wait time in minutes' })
    @IsNumber()
    waitTime: number;
}

class ReturnTimeQueueDto {
    @ApiProperty({ description: 'Current state of return time' })
    @IsEnum(['AVAILABLE', 'UNAVAILABLE'])
    state: string;

    @ApiProperty({ description: 'Start time for return window' })
    @IsDateString()
    returnStart: string;

    @ApiProperty({ description: 'End time for return window' })
    @IsDateString()
    returnEnd: string;
}

class PaidReturnTimeQueueDto extends ReturnTimeQueueDto {
    @ApiProperty({ description: 'Price information' })
    @ValidateNested()
    @Type(() => PriceDto)
    price: PriceDto;
}

class BoardingGroupQueueDto {
    @ApiProperty({ description: 'Status of boarding group allocation' })
    @IsEnum(['AVAILABLE', 'UNAVAILABLE'])
    allocationStatus: string;

    @ApiProperty({ description: 'Start of current boarding group range' })
    @IsNumber()
    currentGroupStart: number;

    @ApiProperty({ description: 'End of current boarding group range' })
    @IsNumber()
    currentGroupEnd: number;

    @ApiProperty({ description: 'Next time boarding groups will be allocated' })
    @IsDateString()
    nextAllocationTime: string;

    @ApiProperty({ description: 'Estimated wait time in minutes' })
    @IsNumber()
    estimatedWait: number;
}

class QueueDto {
    @ApiProperty({ description: 'Standard queue information' })
    @IsOptional()
    @ValidateNested()
    @Type(() => StandardQueueDto)
    STANDBY?: StandardQueueDto;

    @ApiProperty({ description: 'Single rider queue information' })
    @IsOptional()
    @ValidateNested()
    @Type(() => StandardQueueDto)
    SINGLE_RIDER?: StandardQueueDto;

    @ApiProperty({ description: 'Return time queue information' })
    @IsOptional()
    @ValidateNested()
    @Type(() => ReturnTimeQueueDto)
    RETURN_TIME?: ReturnTimeQueueDto;

    @ApiProperty({ description: 'Paid return time queue information' })
    @IsOptional()
    @ValidateNested()
    @Type(() => PaidReturnTimeQueueDto)
    PAID_RETURN_TIME?: PaidReturnTimeQueueDto;

    @ApiProperty({ description: 'Boarding group information' })
    @IsOptional()
    @ValidateNested()
    @Type(() => BoardingGroupQueueDto)
    BOARDING_GROUP?: BoardingGroupQueueDto;

    @ApiProperty({ description: 'Paid standby queue information' })
    @IsOptional()
    @ValidateNested()
    @Type(() => StandardQueueDto)
    PAID_STANDBY?: StandardQueueDto;
}

class ForecastDto {
    @ApiProperty({ description: 'Forecast timestamp' })
    @IsDateString()
    time: string;

    @ApiProperty({ description: 'Predicted wait time in minutes' })
    @IsNumber()
    waitTime: number;

    @ApiProperty({ description: 'Percentage of capacity' })
    @IsNumber()
    percentage: number;
}

class OperatingHoursDto {
    @ApiProperty({ description: 'Type of operating hours' })
    @IsString()
    type: string;

    @ApiProperty({ description: 'End time of operation' })
    @IsDateString()
    endTime: string;

    @ApiProperty({ description: 'Start time of operation' })
    @IsDateString()
    startTime: string;
}

class LiveDataItemDto {
    @ApiProperty({ description: 'Unique identifier of the attraction' })
    @IsString()
    id: string;

    @ApiProperty({ description: 'Name of the attraction' })
    @IsString()
    name: string;

    @ApiProperty({ description: 'Type of entity' })
    @IsEnum(['ATTRACTION'])
    entityType: string;

    @ApiProperty({ description: 'ID of the park this attraction belongs to' })
    @IsString()
    parkId: string;

    @ApiProperty({ description: 'External identifier' })
    @IsString()
    externalId: string;

    @ApiProperty({ description: 'Queue information', required: false })
    @IsOptional()
    @ValidateNested()
    @Type(() => QueueDto)
    queue?: QueueDto;

    @ApiProperty({ description: 'Current operating status' })
    @IsEnum(['OPERATING', 'CLOSED'])
    status: string;

    @ApiProperty({ description: 'Forecast data', required: false })
    @IsOptional()
    @IsArray()
    @ValidateNested({ each: true })
    @Type(() => ForecastDto)
    forecast?: ForecastDto[];

    @ApiProperty({ description: 'Operating hours', required: false })
    @IsOptional()
    @IsArray()
    @ValidateNested({ each: true })
    @Type(() => OperatingHoursDto)
    operatingHours?: OperatingHoursDto[];

    @ApiProperty({ description: 'Last update timestamp' })
    @IsDateString()
    lastUpdated: string;
}

export class LiveDto {
    @ApiProperty({ description: 'Unique identifier of the park' })
    @IsString()
    id: string;

    @ApiProperty({ description: 'Name of the park' })
    @IsString()
    name: string;

    @ApiProperty({ description: 'Type of entity' })
    @IsEnum(['PARK'])
    entityType: string;

    @ApiProperty({ description: 'Timezone of the park' })
    @IsString()
    timezone: string;

    @ApiProperty({ description: 'Live data for attractions', type: [LiveDataItemDto] })
    @IsArray()
    @ValidateNested({ each: true })
    @Type(() => LiveDataItemDto)
    liveData: LiveDataItemDto[];
}