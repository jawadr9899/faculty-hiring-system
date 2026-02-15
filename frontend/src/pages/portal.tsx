
// Mock data to visualize the design better
const jobs = Array.from({ length: 12 }, (_, i) => ({
  id: i,
  title: i % 2 === 0 ? "Senior Frontend Engineer" : "Product Designer",
  company: i % 2 === 0 ? "TechCorp Inc." : "Creative Studio",
  type: i % 3 === 0 ? "Remote" : "Hybrid",
  salary: "$120k - $150k",
  tags: ["React", "Tailwind", "TypeScript"],
  posted: "2d ago"
}));

export default function JobPortal() {
  return (
    <div className="min-h-screen bg-slate-50 text-slate-900 font-sans pb-20">
      {/* --- Hero / Search Section --- */}
      <section className="bg-white border-b border-slate-200 sticky top-0 z-10">
        <div className="max-w-6xl mx-auto px-4 py-8">
          <div className="mb-6">
            <h1 className="text-2xl font-bold text-slate-900">Find your next role</h1>
            <p className="text-slate-500">Browse thousands of job openings with top companies.</p>
          </div>
          
          <div className="relative max-w-2xl">
            {/* Search Icon SVG */}
            <svg 
              className="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" 
              fill="none" 
              stroke="currentColor" 
              viewBox="0 0 24 24"
            >
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
            
            <input 
              type="text" 
              placeholder="Search by job title, keywords, or company..." 
              className="w-full pl-12 pr-4 py-4 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-500/20 focus:border-blue-500 transition-all shadow-sm text-sm sm:text-base"
            />
          </div>
        </div>
      </section>

      {/* --- Job Grid Section --- */}
      <section className="max-w-6xl mx-auto px-4 mt-8">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {jobs.map((job) => (
            <div 
              key={job.id} 
              className="group bg-white rounded-xl border border-slate-200 p-6 hover:shadow-lg hover:border-blue-200 transition-all duration-300 cursor-pointer flex flex-col justify-between"
            >
              {/* Card Header */}
              <div>
                <div className="flex justify-between items-start mb-4">
                  <div className="h-10 w-10 bg-gradient-to-br from-blue-100 to-indigo-100 rounded-lg flex items-center justify-center text-blue-600 font-bold text-sm">
                    {job.company.charAt(0)}
                  </div>
                  <span className="text-xs font-medium text-slate-400 bg-slate-50 px-2 py-1 rounded-full border border-slate-100">
                    {job.posted}
                  </span>
                </div>

                <h3 className="font-bold text-lg text-slate-900 mb-1 group-hover:text-blue-600 transition-colors">
                  {job.title}
                </h3>
                <p className="text-sm text-slate-500 mb-4">{job.company}</p>

                {/* Tags / Details */}
                <div className="flex flex-wrap gap-2 mb-6">
                  <span className="inline-flex items-center px-2.5 py-0.5 rounded-md text-xs font-medium bg-blue-50 text-blue-700">
                    {job.type}
                  </span>
                  <span className="inline-flex items-center px-2.5 py-0.5 rounded-md text-xs font-medium bg-green-50 text-green-700">
                    {job.salary}
                  </span>
                </div>
              </div>

              {/* Card Footer */}
              <div className="pt-4 border-t border-slate-100 flex items-center justify-between">
                <div className="flex -space-x-2 overflow-hidden">
                   {/* Fake avatars for "Applicants" social proof */}
                   {[1,2,3].map(a => (
                     <div key={a} className="inline-block h-6 w-6 rounded-full ring-2 ring-white bg-slate-200" />
                   ))}
                   <span className="text-xs text-slate-400 pl-3 self-center">+12 applicants</span>
                </div>
                
                <button className="text-sm font-semibold text-slate-900 group-hover:underline decoration-blue-500 underline-offset-4">
                  Apply Now &rarr;
                </button>
              </div>
            </div>
          ))}
        </div>
      </section>
    </div>
  )
}