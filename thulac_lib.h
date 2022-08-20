int init(const char* model, const char* dict, int ret_size, int t2s, int just_seg);
void deinit();
char *getResult();
void freeResult();
int seg(const char *in);