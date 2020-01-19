import { Component, OnInit } from '@angular/core';
import {ApiErrorResponse, ResponseStatus} from '../../interfaces/api';
import {AuthService} from '../../services/auth.service';
import {Router} from '@angular/router';
import {StorageService} from '../../services/storage.service';

@Component({
  selector: 'app-logout',
  templateUrl: './logout.component.html',
  styleUrls: ['./logout.component.scss']
})
export class LogoutComponent implements OnInit {
  constructor(
    private readonly auth: AuthService,
    private readonly router: Router,
    private readonly storage: StorageService
  ) { }

  public logout(): void {
    this.auth.logout()
      .then(res => {
        if (res.status === ResponseStatus.Ok) {
          this.storage.saveSession(null);
          return this.router.navigate(['/login']);
        } else {
          return Promise.reject((res as ApiErrorResponse).error_detail);
        }
      }).catch(err => console.log(err));
  }

  ngOnInit() {
  }
}
