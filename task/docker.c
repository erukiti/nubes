#include <stdio.h>
#include "mruby.h"
#include "_cgo_export.h"

mrb_value docker_build(mrb_state *mrb, mrb_value value) {
	char *tag;
	char *baseDir;

	mrb_get_args(mrb, "zz", &tag, &baseDir);
	dockerBuild(tag, baseDir);
	return value;
}

mrb_value docker_run(mrb_state *mrb, mrb_value value) {
	char *image;
	char *id;

	mrb_get_args(mrb, "z", &image);
	id = dockerRun(image);
	if (id != NULL) {
		return mrb_str_new_cstr(mrb, id);
	} else {
		return mrb_nil_value();
	}
}

void docker_init(mrb_state *mrb) {
	struct RClass *d;
	d = mrb_define_module(mrb, "Docker");
	mrb_define_class_method(mrb, d, "build", docker_build, MRB_ARGS_REQ(2));
	mrb_define_class_method(mrb, d, "run", docker_run, MRB_ARGS_REQ(1));
}
