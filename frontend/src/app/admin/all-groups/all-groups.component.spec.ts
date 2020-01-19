import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AllGroupsComponent } from './all-groups.component';

describe('GroupsListComponent', () => {
  let component: AllGroupsComponent;
  let fixture: ComponentFixture<AllGroupsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AllGroupsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AllGroupsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
