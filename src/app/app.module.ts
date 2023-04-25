import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NZ_I18N } from 'ng-zorro-antd/i18n';
import { zh_CN } from 'ng-zorro-antd/i18n';
import { registerLocaleData } from '@angular/common';
import zh from '@angular/common/locales/zh';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { HistorysComponent } from './history/historys/historys.component';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzDividerModule } from 'ng-zorro-antd/divider';
import { NzSpaceModule } from 'ng-zorro-antd/space';
import { NzSpinModule } from 'ng-zorro-antd/spin';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzMessageModule } from 'ng-zorro-antd/message';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzInputModule } from 'ng-zorro-antd/input';
import { JobsComponent } from './job/jobs/jobs.component';
import { JobEditComponent } from './job/job-edit/job-edit.component';
import { NzInputNumberModule } from 'ng-zorro-antd/input-number';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzTagModule } from 'ng-zorro-antd/tag';
import {APP_BASE_HREF } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import {BaseModule} from "./base/base.module";
import {NzCardModule} from "ng-zorro-antd/card";
import { NzCollapseModule } from 'ng-zorro-antd/collapse'; 
 
registerLocaleData(zh);

@NgModule({
  declarations: [
    AppComponent,
    HistorysComponent,
    JobsComponent,
    JobEditComponent
  ],
  imports: [
    BrowserModule,
    NzButtonModule,
    NzDividerModule,
    NzCollapseModule,
    NzSpaceModule,
    NzTableModule ,
    NzSpaceModule,
    NzMessageModule,
    NzInputNumberModule,
    NzTagModule,
    NzFormModule,
    NzCardModule,
    ReactiveFormsModule,
    NzIconModule,
    AppRoutingModule,
    BaseModule,
    NzSelectModule ,
    NzSpinModule,
    NzInputModule,
    NzModalModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule
  ],
  providers: [
    { provide: APP_BASE_HREF, useValue: '/app/history/' },
    { provide: NZ_I18N, useValue: zh_CN }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
