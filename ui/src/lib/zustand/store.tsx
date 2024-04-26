import { create } from 'zustand'
import { data } from "./employeeData"

export interface IEmployee {
    id: number
    first_name: string
    last_name: string
    email: string
    position: string
    department: string
}
interface EmployeeState {
    employees: IEmployee[]
    isEditing: boolean
    currentEdit: IEmployee | undefined
    create: (data: IEmployee) => void
    update: (data: IEmployee[]) => void
    delete: (data: number) => void
    setIsEditing: ({ data, currentEdit }: { data: IEmployee, currentEdit: IEmployee | undefined }) => void
}

const useEmployeeStore = create<EmployeeState>()((set) => ({
    employees: data as IEmployee[],
    isEditing: false,
    currentEdit: undefined,
    create: (by) => set((state) => ({ employees: [...state.employees, by] })),
    update: (by) => set((state) => {
        let newEmployees: IEmployee[] = [...state.employees];

        by.forEach(a => {
            const empID = state.employees.findIndex(x => x.id == a.id)
            newEmployees[empID] = a
        })

        return { employees: newEmployees }
    }),
    delete: (data) => set(state => {
        return { employees: state.employees.filter(a => a.id != data) }
    }),
    setIsEditing: ({ data, currentEdit }) => set(() => {
        return { isEditing: data, currentEdit: currentEdit, }
    }),
}))
export default useEmployeeStore
