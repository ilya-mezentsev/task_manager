import { Injectable } from '@angular/core';
import {Router} from '@angular/router';
import {SessionResponse} from '../interfaces/api';

@Injectable({
  providedIn: 'root'
})
export class RoleNavigationService {
  private readonly statusError: string = 'error';

  constructor(
    private readonly router: Router
  ) { }

  public navigateUser(session: SessionResponse): void {
    if (session.status === this.statusError) {
      return;
    }

    this.router.navigate([session.data.role])
      .catch(err => console.log(`Navigation error: ${err}`));
  }
}
