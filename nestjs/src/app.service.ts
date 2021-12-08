import { Injectable } from "@nestjs/common";

@Injectable()
export class AppService {
  isOnline(): { server: string } {
    return { server: "online" };
  }
}
