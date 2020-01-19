import { Injectable } from '@angular/core';
import {Router} from '@angular/router';
import {SessionResponse, UserRole} from '../interfaces/api';

@Injectable({
  providedIn: 'root'
})
export class RoleNavigationService {
  constructor(
    private readonly router: Router
  ) { }

  public navigateUser(session: SessionResponse): void {
    this.router.navigate([RoleNavigationService.getNavigationSegmentBy(session.data.role)])
      .catch(err => console.log(`Navigation error: ${err}`));
  }

  private static getNavigationSegmentBy(userRole: UserRole): string {
    switch (userRole) {
      case 'admin':
        return 'admin';
      case 'group_lead':
      case 'group_worker':
        return 'group';
      default:
        return 'not-found';
    }
  }
}
