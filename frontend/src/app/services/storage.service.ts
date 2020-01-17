import { Injectable } from '@angular/core';
import {UserSession} from '../interfaces/api';

@Injectable({
  providedIn: 'root'
})
export class StorageService {
  private readonly storage = localStorage;
  private readonly itemsKeys: {[key: string]: string} = {
    session: 'user_session'
  };

  public saveSession(session: UserSession): void {
    this.save(this.itemsKeys.session, session);
  }

  public getSession(): UserSession {
    return JSON.parse(this.get(this.itemsKeys.session));
  }

  private save(key: string, value: any): void {
    this.storage.setItem(key, JSON.stringify(value));
  }

  private get(key: string): string {
    return this.storage.getItem(key);
  }
}
