import { Controller, Redirect, Get } from '@nestjs/common';

@Controller()
export class RootController {
    @Get()
    @Redirect('/docs', 302)
    redirectToDocs() {
        // This method will redirect to /docs
    }
}
