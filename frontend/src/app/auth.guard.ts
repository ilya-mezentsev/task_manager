import { Injectable } from '@angular/core';
import {CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, Router, UrlTree} from '@angular/router';
import {StorageService} from './services/storage.service';
import {ApiErrorResponse, SessionResponse, UserSession} from './interfaces/api';
import {AuthService} from './services/auth.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  private readonly loginUrlTree: UrlTree;

  constructor(
    private readonly router: Router,
    private readonly storage: StorageService,
    private readonly authService: AuthService
  ) {
    this.loginUrlTree = this.router.parseUrl('login');
  }

  async canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): Promise<boolean | UrlTree> {
    const roles: string[] = next.data.roles as Array<string>;
    const session: UserSession = this.storage.getSession();
    if (session !== null) {
      return roles.includes(session.role)
        ? true
        : this.loginUrlTree;
    }

    return await this.loadAndCheckSession(roles);
  }

  private async loadAndCheckSession(roles: string[]): Promise<boolean | UrlTree> {
    try {
      const sessionResponse = await this.authService.getSession();
      if (sessionResponse.status === 'error') {
        console.log(`an error while getting session: ${(sessionResponse as ApiErrorResponse).error_detail}`);
        return false;
      } else {
        const session = (sessionResponse as SessionResponse).data;
        this.storage.saveSession(session);
        return roles.includes(session.role)
          ? true
          : this.loginUrlTree;
      }
    } catch (e) {
      console.log(`an error while getting session: ${e}`);
      return this.loginUrlTree;
    }
  }
}
