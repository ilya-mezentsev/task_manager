import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-group-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.scss']
})
export class NavigationComponent implements OnInit {
  @Input() public readonly isGroupLead: boolean;

  constructor() { }

  ngOnInit() {
  }
}
