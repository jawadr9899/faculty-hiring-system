import React from 'react';
import { ArrowUpRight } from 'lucide-react';
import { type StatMetric } from '@/types/types';

const StatsCard: React.FC<StatMetric> = ({ title, value, trend, icon: Icon, color }) => (
  <div className="bg-slate-900 border border-slate-800 rounded-xl p-5 hover:border-slate-700 transition-colors">
    <div className="flex justify-between items-start mb-4">
      <div className={`p-2 rounded-lg bg-slate-800 ${color}`}>
        <Icon size={20} />
      </div>
      {trend && (
        <span className="flex items-center text-xs font-medium text-emerald-400 bg-emerald-400/10 px-2 py-0.5 rounded-full">
          <ArrowUpRight size={12} className="mr-1" /> {trend}
        </span>
      )}
    </div>
    <div className="text-2xl   font-bold text-white mb-1">{value}</div>
    <div className="text-sm text-slate-500">{title}</div>
  </div>
);

export default StatsCard;