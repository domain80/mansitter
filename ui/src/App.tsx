import { useState } from 'react'
import './App.css'
import { SideNav } from './components/side-nav'
import { Bell, Search } from 'lucide-react'
import { Button } from './components/ui/button'
import { Input } from './components/ui/input'
import { DataTableDemo } from './components/custom/DataTable'

function App() {

  return (
    <div className="flex" >
      <SideNav />
      <main className='flex-1 w-full px-4 py-4'>
        <header className='flex justify-between'>
          <div>

            <h1 className="text-2xl font-bold">
              All Employees
            </h1>
            <p className='text-sm text-gray-500'>All employeee information</p>

          </div>
          <div>
            <Button variant={"outline"} className=''><Bell className='w-5' /></Button>
          </div>
        </header>


        <section>
          <DataTableDemo />
        </section>
      </main>

    </div>
  )
}

export default App
