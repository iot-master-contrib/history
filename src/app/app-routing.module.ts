import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router'; 
import { HistorysComponent } from './history/historys/historys.component';
import { JobEditComponent } from './history/job/job-edit/job-edit.component';
import { JobsComponent } from './history/job/jobs/jobs.component';
import { PageNotFoundComponent } from './base/page-not-found/page-not-found.component';
const routes: Routes = [
  
 { path: '', pathMatch: 'full', redirectTo: 'history' },
    { path: 'history', component: HistorysComponent },
    { path: 'job', component: JobsComponent },
    { path: 'job/edit/:id', component: JobEditComponent },
    { path: 'job/create', component: JobEditComponent },
    { path: '**', component: PageNotFoundComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
