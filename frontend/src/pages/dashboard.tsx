import React, { useState } from 'react';
import { Briefcase, Users, LayoutDashboard, Bell, Menu } from 'lucide-react';
import Sidebar from "@/components/sidebar.dashboard"
import StatsCard from "@/components/stats-card.dashboard";
import { PerformanceScoreChart, CompositeWeightChart } from "@/components/analytics-charts.dashboard";
import JobsView from "@/components/job-views.dashboard";
import CandidatesView from "@/components/candidates-view.dashboard"
import { type Tab } from '@/types/types';

const AdminDashboard: React.FC = () => {
  const [activeTab, setActiveTab] = useState<Tab>('overview');
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  // --- Render the Overview Tab Content ---
  const renderOverview = () => (
    <div className="space-y-6 animate-fadeIn">
      {/* Top Stats */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        <StatsCard title="Total Applications" value="1,293" trend="+12%" icon={Users} color="text-blue-500" />
        <StatsCard title="Active Jobs" value="24" trend="+4" icon={Briefcase} color="text-purple-500" />
        <StatsCard title="Avg Response" value="2 Days" icon={Bell} color="text-orange-500" />
      </div>

      {/* Charts Grid */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <div className="lg:col-span-2">
          <PerformanceScoreChart />
        </div>
        <div>
          <CompositeWeightChart />
        </div>
      </div>
    </div>
  );

  return (
    <div className="flex h-screen bg-slate-950 text-slate-200 font-sans selection:bg-blue-500 selection:text-white overflow-hidden">
      
      {/* 1. Sidebar Component */}
      <Sidebar activeTab={activeTab} setActiveTab={setActiveTab} />

      {/* 2. Main Content Area */}
      <div className="flex-1 flex flex-col min-w-0">
        
        {/* Header */}
        <header className="h-16 bg-slate-900 border-b border-slate-800 flex items-center justify-between px-6 z-20 flex-shrink-0">
          <div className="flex items-center gap-4">
            <button className="md:hidden text-slate-400" onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}>
              <Menu size={24} />
            </button>
            <h2 className="text-lg font-bold text-white capitalize">{activeTab}</h2>
          </div>
          
          <div className="flex items-center gap-4">
            <button className="relative p-2 text-slate-400 hover:text-white hover:bg-slate-800 rounded-full transition-colors">
              <Bell size={20} />
              <span className="absolute top-2 right-2 h-2 w-2 bg-red-500 rounded-full border border-slate-900"></span>
            </button>
            <div className="h-8 w-8 rounded-full bg-gradient-to-tr from-blue-500 to-purple-500 border border-slate-700"></div>
          </div>
        </header>

        {/* Scrollable Canvas */}
        <main className="flex-1 overflow-y-auto p-4 md:p-8 bg-slate-950 relative">
          <div className="max-w-6xl mx-auto pb-12">
            {activeTab === 'overview' && renderOverview()}
            {activeTab === 'jobs' && <JobsView />}
            {activeTab === 'candidates' && <CandidatesView />}
          </div>
        </main>
      </div>
    </div>
  );
};

export default AdminDashboard;