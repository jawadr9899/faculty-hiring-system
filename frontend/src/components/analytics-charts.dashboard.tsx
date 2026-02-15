import { 
  AreaChart, 
  Area, 
  XAxis, 
  YAxis, 
  CartesianGrid, 
  Tooltip, 
  ResponsiveContainer,
  PieChart,
  Pie,
  Sector
} from 'recharts';

// Mock data based on your new fields
const PERFORMANCE_DATA = [
  { category: 'Academic', score: 85, rank: 2 },
  { category: 'Research', score: 92, rank: 1 },
  { category: 'Teaching', score: 78, rank: 5 },
  { category: 'Industrial', score: 65, rank: 8 },
  { category: 'Admin', score: 70, rank: 6 },
  { category: 'Salary', score: 88, rank: 3 },
];

const SCORE_DISTRIBUTION = [
  { name: 'Research', value: 35, fill: '#3b82f6' },
  { name: 'Academic', value: 25, fill: '#8b5cf6' },
  { name: 'Teaching', value: 20, fill: '#10b981' },
  { name: 'Others', value: 20, fill: '#f59e0b' },
];

const renderActiveShape = (props: any) => {
  const { cx, cy, innerRadius, outerRadius, startAngle, endAngle, payload } = props;
  return (
    <Sector
      cx={cx}
      cy={cy}
      innerRadius={innerRadius}
      outerRadius={outerRadius}
      startAngle={startAngle}
      endAngle={endAngle}
      fill={payload.fill}
      stroke={payload.fill}
    />
  );
};

export const PerformanceScoreChart = () => (
  <div className="bg-slate-900 border border-slate-800 rounded-xl p-6 h-full">
    <div className="flex justify-between items-start mb-6">
      <div>
        <h3 className="text-lg font-semibold text-white">Score Analytics</h3>
        <p className="text-xs text-slate-400">Composite Rank: #1 (Senior Research Level)</p>
      </div>
    </div>
    <div className="h-[300px] w-full">
      <ResponsiveContainer width="100%" height="100%">
        <AreaChart data={PERFORMANCE_DATA}>
          <defs>
            <linearGradient id="scoreGradient" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor="#10b981" stopOpacity={0.3}/>
              <stop offset="95%" stopColor="#10b981" stopOpacity={0}/>
            </linearGradient>
          </defs>
          <CartesianGrid strokeDasharray="3 3" stroke="#1e293b" vertical={false} />
          <XAxis dataKey="category" stroke="#64748b" fontSize={12} tickLine={false} axisLine={false} />
          <YAxis stroke="#64748b" fontSize={12} tickLine={false} axisLine={false} domain={[0, 100]} />
          <Tooltip 
            contentStyle={{ backgroundColor: '#0f172a', borderColor: '#334155', borderRadius: '8px', color: '#fff' }}
          />
          <Area 
            type="monotone" 
            dataKey="score" 
            stroke="#10b981" 
            fillOpacity={1} 
            fill="url(#scoreGradient)" 
            strokeWidth={3} 
          />
        </AreaChart>
      </ResponsiveContainer>
    </div>
  </div>
);

export const CompositeWeightChart = () => (
  <div className="bg-slate-900 border border-slate-800 rounded-xl p-6 h-full">
    <h3 className="text-lg font-semibold text-white mb-2">Composite Breakdown</h3>
    <p className="text-xs text-slate-400 mb-6">Degree: PhD | Exp: 12 Years</p>
    <div className="h-[250px] w-full relative">
      <ResponsiveContainer width="100%" height="100%">
        <PieChart>
          <Pie
            data={SCORE_DISTRIBUTION}
            cx="50%"
            cy="50%"
            innerRadius={60}
            outerRadius={80}
            paddingAngle={8}
            dataKey="value"
            shape={renderActiveShape}
            stroke="none"
          />
          <Tooltip />
        </PieChart>
      </ResponsiveContainer>
      
      <div className="flex flex-wrap justify-center gap-3 mt-4">
        {SCORE_DISTRIBUTION.map((item, idx) => (
          <div key={idx} className="flex items-center gap-1.5 text-xs text-slate-400">
            <span className="w-2 h-2 rounded-full" style={{ backgroundColor: item.fill }}></span>
            {item.name} ({item.value}%)
          </div>
        ))}
      </div>
    </div>
  </div>
);