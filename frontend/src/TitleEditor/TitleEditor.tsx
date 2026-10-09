import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import styles from './TitleEditor.module.css'

const titleSchema = z.object({ title: z.string().trim().min(1, 'A name is required.').max(100) })
export type TitleFormValues = z.infer<typeof titleSchema>

type TitleEditorProps = {
  label: string
  placeholder: string
  busy: boolean
  onSubmit: (values: TitleFormValues) => void
  onCancel: () => void
}

const TitleEditor = ({ label, placeholder, busy, onSubmit, onCancel }: TitleEditorProps) => {
  const { register, handleSubmit, formState: { errors } } = useForm<TitleFormValues>({
    resolver: zodResolver(titleSchema),
    defaultValues: { title: '' },
  })

  return (
    <form className={styles.form} onSubmit={handleSubmit(onSubmit)}>
      <label>
        {label}
        <input autoFocus placeholder={placeholder} {...register('title')} />
        {errors.title && <small>{errors.title.message}</small>}
      </label>
      <div className={styles.modalActions}>
        <button type="button" className={styles.cancelButton} onClick={onCancel}>Cancel</button>
        <button className={styles.primaryButton} disabled={busy}>{busy ? 'Saving…' : 'Create'}</button>
      </div>
    </form>
  )
}

export default TitleEditor
