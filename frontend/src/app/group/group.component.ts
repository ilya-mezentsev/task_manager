import { Component, OnInit } from '@angular/core';
import {StorageService} from '../services/storage.service';
import {UserSession} from '../interfaces/api';

@Component({
  selector: 'app-group',
  templateUrl: './group.component.html',
  styleUrls: ['./group.component.scss']
})
export class GroupComponent implements OnInit {
  public isGroupLead: boolean = false;

  constructor(
    private readonly storage: StorageService
  ) { }

  ngOnInit() {
    const session: UserSession = this.storage.getSession();
    if (session != null) {
      this.isGroupLead = session.role === 'group_lead';
    }
  }
}
