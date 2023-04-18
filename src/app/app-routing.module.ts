import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { HistorysComponent } from './history/historys/historys.component';
const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: '/alarm/alarms' },
  { path: 'history/historys', component: HistorysComponent },
  { path: '**', component: PageNotFoundComponent },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
