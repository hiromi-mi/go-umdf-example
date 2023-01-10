static void __guard_check_icall_dummy(void) {}
__asm__(
    "__guard_dispatch_icall_dummy:\n"
    "    jmp *%rax\n"
);
extern void *__guard_dispatch_icall_dummy;
////  __asm__(".section .00cfg\n");
// //__attribute__(( section (".data") ))
void *__guard_check_icall_fptr = &__guard_check_icall_dummy;
// //__attribute__(( section (".data") ))
void *__guard_dispatch_icall_fptr = &__guard_dispatch_icall_dummy; 
void *__guard_xfg_dispatch_icall_fptr = &__guard_dispatch_icall_dummy;
void *_guard_xfg_check_icall_fptr = &__guard_dispatch_icall_dummy;
void *_castguard_check_failure_os_handled_fptr =  &__guard_dispatch_icall_dummy;
void *_guard_xfg_table_dispatch_icall_fptr =  &__guard_dispatch_icall_dummy;