import type { ReactNode } from 'react'
import styles from './Icon.module.css'

type IconProps = { children: ReactNode }

const Icon = ({ children }: IconProps) => (
  <span className={styles.icon} aria-hidden="true">{children}</span>
)

export default Icon
