import { ApiProperty } from "@nestjs/swagger";
import { Type } from "class-transformer";
import { ValidateNested, IsArray } from "class-validator";

class ParkDto {
    @ApiProperty({ example: '12345678-1234-1234-1234-123456789012', description: 'Park id' })
    id: string;

    @ApiProperty({ example: 'Magic Kingdom', description: 'Name of the park' })
    name: string;
}

export class DestinationDto {
  @ApiProperty({ example: '12345678-1234-1234-1234-123456789012', description: 'Destination id' })
  id: string;

  @ApiProperty({ example: 'Walt Disney World Resort', description: 'Name of the destination which may have multiple parks' })
  name: string;

  @ApiProperty({ example: 'waltdisneyworldresort', description: 'Short name identifier for the destination' })
  slug: string;

  @ApiProperty({ type: [ParkDto], description: 'List of parks within the destination' })
  @IsArray()
  @ValidateNested({ each: true })
  @Type(() => ParkDto)
  parks: ParkDto[];
}